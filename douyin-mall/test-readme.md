测试用例：
1. 测试支付服务：PaymentService/Charge:
{
  "amount": 100.00,
  "credit_card": {
    "credit_card_number": "4111111111111111",
    "credit_card_cvv": 123,
    "credit_card_expiration_year": 2025,
    "credit_card_expiration_month": 12
  },
  "order_id": "ORDER_001",
  "user_id": 1001
}
2. 测试交易状态：PaymentService/GetTransactionStatus:
{
    "transaction_id": "交易id"
}
