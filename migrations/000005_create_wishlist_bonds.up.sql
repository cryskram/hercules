CREATE TABLE wishlist_bonds (
    wishlist_id UUID NOT NULL,
    bond_isin CHAR(12) NOT NULL,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (wishlist_id, bond_isin),
    CONSTRAINT fk_wishlist FOREIGN KEY (wishlist_id) REFERENCES wishlists(id) ON DELETE CASCADE,
    CONSTRAINT fk_bond FOREIGN KEY (bond_isin) REFERENCES bonds(isin) ON DELETE CASCADE
);