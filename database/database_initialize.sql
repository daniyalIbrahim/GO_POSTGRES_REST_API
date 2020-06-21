
DROP TABLE customers,suppliers,consignments;


ALTER TABLE consignments
ADD CONSTRAINT con_customer_fk
FOREIGN KEY (customer_id)
REFERENCES customers(customer_id)
ON DELETE CASCADE;

ALTER TABLE consignments
ADD CONSTRAINT con_supplier_fk
FOREIGN KEY (supplier_id)
REFERENCES suppliers(supplier_id)
ON DELETE CASCADE;


ALTER TABLE van
ADD CONSTRAINT van_supplier_fk
FOREIGN KEY (supplier_id)
REFERENCES suppliers(supplier_id)
ON DELETE CASCADE;