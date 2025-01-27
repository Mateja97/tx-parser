# Tx-Parser

A simple Ethereum blockchain parser that allows you to query transactions for subscribed addresses.

---

## **Run the Project**

To start the project, use the following command:

```bash
docker-compose up
```

---

## **API Endpoints**

### **1. Subscribe**
Subscribe to a specific Ethereum address to monitor its transactions.

- **Request:**
    ```json
    {
        "address": "YOUR_ADDRESS"
    }
    ```

- **Response:**
  A message indicating whether the subscription was successful or not.

---

### **2. Get Transactions**
Retrieve all inbound and outbound transactions for a subscribed Ethereum address.

- **Request:**
    ```json
    {
        "address": "YOUR_ADDRESS"
    }
    ```

- **Response:**
  A list of all transactions related to the given address.

---

### **3. Get Current Block**
Fetch the latest Ethereum block number.

- **Request:**
    ```json
    {}
    ```

- **Response:**
  The current block number.

---
