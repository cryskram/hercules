CREATE INDEX idx_bond_name_trgm ON bonds USING gin (bond_name gin_trgm_ops);
CREATE INDEX idx_brand_name_trgm ON bonds USING gin (brand_name gin_trgm_ops);
CREATE INDEX idx_isin ON bonds(isin);