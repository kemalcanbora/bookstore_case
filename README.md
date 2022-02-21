## Bookstore Demo

#### API Collection
 You can find the collection of books n the [`build`] folder.
#### How to build dockerfile ?
- docker-compose build

#### How to run it ?
- `cd build`
- `docker-compose up`

Note: if you just want to run mongodb through docker and want to run the API through the localhost
- `nano .env` 
- change `mongodb://mongodb:27017` to `mongodb://localhost:27017`

#### Endpoints
 - `/api/user/login` : You can login with this endpoint 
   - Example: ` {"email":"customer@gmail.com", "password":"123456"}`
 - `/api/user/registration` :  You can register with this endpoint 
   - Example: `{"email":"customer@gmail.com", "password":"123456"}`
 - `/api/stock/add` :  Add item to stock.
 - `/api/stock/update`: Update item in stock.
 - `/api/stock/delete` : Delete item in stock.
 - `/api/order/buy` :  You can order a book.


  