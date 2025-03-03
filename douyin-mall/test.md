# User

## 注册示例

 {

​    "email": "222@example.com",

​    "password": "123456",

​    "confirm_password": "123456"

}

## 登录

   {
       "email": "222@example.com",
       "password": "123456"
   }

## 获取用户信息 

{
    "user_id": 4
}

# product

## ListProduct

{

  "categoryName": "",

  "page": 1,

  "pageSize": 10

}

## getproduct

{

  "id": 1

}

## SearchListProduct

{

 "query": "衣物"

}

## CreateProduct

{

  	"categories": [

​    		"Test Category","电子"

  	],

  	"description": "This is a test product4.",

  	"name": "Product4",

 	 "picture": "http://example.com/product3.jpg",

  	"price": 20.00

}

## DeleteProducts

{

  "id": 4

}

## UpdateProducts

{

  "categories": [

​    "Test Category","衣物"

  ],

  "description": "This is a test product3.",

  "id": 3,

  "name": "Product1",

  "picture": "http://example.com/product3.jpg",

  "price": 9.99

}

# Cart

## AddItem

#多加

{

  "item": {

​    "product_id": 4,

​    "quantity": 2

  },

  "user_id": 4

}

## GetCart

{

  "user_id": 4

}

## EmptyCart

{

  "user_id": 1

}

#重新添加一个购物车

# Order

## PlaceOrder

{

  "address": {

​    "city": "Test City",

​    "country": "USA",

​    "state": "TX",

​    "street_address": "123 Test St",

​    "zip_code": 12345

  },

  "email": "222@example.com",

  "order_items": [

​    {

​      "item": {

​        "product_id": 4,

​        "quantity": 2

​      },

​      "cost": 40

​    },

​    {

​      "item": {

​        "product_id": 5,

​        "quantity": 1

​      },

​      "cost": 15

​    }

  ],

  "user_currency": "USD",

  "user_id": 4

}

## ListOrder

{

  "user_id": 4

}

## MarkOrderPaid

#结算后查询

{

  "order_id": "",

  "user_id": 4

}

# checkout

## checkout

{

  "address": {

​    "city": "Test City",

​    "country": "USA",

​    "state": "TX",

​    "street_address": "123 Test St",

​    "zip_code": "12345"

  },

  "credit_card": {

​    "credit_card_cvv": 123,

​    "credit_card_expiration_month": 12,

​    "credit_card_expiration_year": 2025,

​    "credit_card_number": "4111111111111132"

  },

  "email": "222@example.com",

  "firstname": "wu",

  "lastname": "ZHANG",

  "user_id": 4

}

# payment

## charge

{
  "amount": 100.00,
  "credit_card": {
 "credit_card_number": "4111111111111111",
 "credit_card_cvv": 123,
 "credit_card_expiration_year": 2025,
 "credit_card_expiration_month": 12
  },
  "order_id": "",
  "user_id": 1
}

## gettranslationstaus

{

  "transaction_id": "exercitation Excepteur laborum minim"

}