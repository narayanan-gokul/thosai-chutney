CREATE OR REPLACE FUNCTION check_order_less_than_max()
RETURNS TRIGGER AS $check_order_less_than_max$
BEGIN
	IF NEW.quantity > (SELECT max_cap FROM item WHERE item_id = NEW.item_id) THEN
		RAISE EXCEPTION 'Order quantity % exceeds max allowed for item %', NEW.quantity, NEW.item_id;
	END IF;
	RETURN NEW;
END;
$check_order_less_than_max$ LANGUAGE plpgsql;
