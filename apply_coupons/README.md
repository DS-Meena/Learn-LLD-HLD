# Apply coupons on Shopping card Products

Problem statement:
"Given shopping cart with products and coupons and calculate the net price after applying coupons on products.
1. N% off that is 10% off for all the individual products.
2. P% off on next item
3. D% off on Nth item of type T
Sequentially wants to apply all the coupons on the card and get the Total amount.
"

## Solution

Let's try to solve this problem using decorator pattern. We will use decoratorCoupons, that will applied over the base products.

## Components

1. Product - will be wrapped in coupon decorator. 
2. Coupon - applied over products in list, can have different rules to calculate price.
3. Shopping Cart - has products list.

`product` interface that has a price method.

`couponDecorator` interface implements product, with its own type of discount method. and has product.

Product - (price())
couponDecorator - (discount, Product, price())
Cart - (price(), list)

## Reference

[You tube](https://www.youtube.com/watch?v=EfQesfKZ3Jw&list=PL6W8uoQQ2c61X_9e6Net0WdYZidm7zooW&index=39)