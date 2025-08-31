SELECT 
	postcode, item_id, SUM(quantity)
FROM 
	item
	JOIN shipment ON item.item_id = shipment.item_id 
	JOIN supplier ON supplier.supp_id = shipment.supp_id
GROUP BY
	postcode, item_id;
