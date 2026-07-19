-- name: GetUser :one
SELECT * FROM users WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO users(firstname,lastname,email,password,line1,line2,city,state,zipcode) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9);

-- name: CreateOrder :exec
INSERT INTO orders(userid,itemName,quantity,status,instruction) VALUES ($1,$2,$3,$4,$5);