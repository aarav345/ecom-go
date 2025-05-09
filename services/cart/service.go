package cart

import (
	"fmt"

	"github.com/aarav345/ecom-go/types"
)

func getCardItemIDs(items []types.CartItem) ([]int, error) {
	productIDs := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductID)
		}

		productIDs[i] = item.ProductID
	}
	return productIDs, nil
}

func (h *Handler) createOrder(ps []types.ProductWithInventory, items []types.CartItem, userID int, updateProductFields bool) (int, float64, error) {

	productMap := make(map[int]types.ProductWithInventory)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	// check if all the products are actually in stock
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}
	// calculate the total price

	totalPrice := calculateTotalPrice(items, productMap)

	// reduce quantity of the products in our db
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity
		h.productStore.UpdateProduct(product, updateProductFields)
	}

	// create the order
	orderID, err := h.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "some address",
	})

	if err != nil {
		return 0, 0, nil
	}

	// create order items
	for _, item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}

	return orderID, totalPrice, nil
}

func calculateTotalPrice(cartItem []types.CartItem, products map[int]types.ProductWithInventory) float64 {
	var totalPrice float64
	for _, item := range cartItem {
		product := products[item.ProductID]
		totalPrice += product.Price * float64(item.Quantity)
	}

	return totalPrice
}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[int]types.ProductWithInventory) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}

	return nil
}
