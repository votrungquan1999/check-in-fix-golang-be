package purchases

import "checkinfix.com/models"

func CalculateCosts(
	purchaseProducts []models.PurchaseProducts,
	purchaseDiscount float64,
	shippingFee float64,
) (grandTotal float64, totalDiscount float64) {
	for _, purchaseProduct := range purchaseProducts {
		subTotal := float64(purchaseProduct.Quantity) * purchaseProduct.UnitCost
		grandTotal += subTotal - purchaseProduct.Discount
		totalDiscount += purchaseProduct.Discount
	}

	grandTotal -= purchaseDiscount
	totalDiscount += purchaseDiscount

	grandTotal += shippingFee

	return
}
