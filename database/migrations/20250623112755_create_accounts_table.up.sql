CREATE TABLE accounts (
	id SERIAL PRIMARY KEY NOT NULL,
	user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	name VARCHAR(100) NOT NULL,
	account_type VARCHAR(20) NOT NULL,
	balance DECIMAL(15,2) default 0.00,
	currency VARCHAR(3) DEFAULT 'IDR',
	is_active BOOLEAN DEFAULT TRUE,
	created_at timestamp WITH TIME ZONE DEFAULT NOW(),
	updated_at timestamp WITH TIME ZONE DEFAULT NOW(),
	deleted_at timestamp WITH TIME ZONE NULL
);

CREATE INDEX idx_accunts_type on accounts(account_type);	