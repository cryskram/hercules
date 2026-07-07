CREATE INDEX idx_bonds_brand_name ON bonds(brand_name);
CREATE INDEX idx_bonds_sector ON bonds(sector);
CREATE INDEX idx_bonds_rating ON bonds(rating);
CREATE INDEX idx_bonds_yield ON bonds(yield_pct DESC);
CREATE INDEX idx_bonds_maturity ON bonds(maturity_date);
CREATE INDEX idx_bonds_coupon_type ON bonds(coupon_type);
CREATE INDEX idx_bonds_payout_frequency ON bonds(payout_frequency);
CREATE INDEX idx_bonds_nature ON bonds(nature);
CREATE INDEX idx_bonds_is_active ON bonds(is_active);