# Hacer pagos Openpay

Openpay nos permite crear clientes, asociarles tarjetas y realizar cargos 
a estas tarjetas

## Running the sample

1. Correr el proyecto:

~~~
go run server.go
~~~

2. Ir a: [http://localhost:4242/register.html](http://localhost:4242/register.html)

~~~
Registramos a nuestro cliente y agregamos su ID al archivo server.go
~~~

3. Ir a: [http://localhost:4242/card.html](http://localhost:4242/card.html)

~~~
Capturamos la tarjeta del cliente para asociarla dentro de la plataforma
y agregamos el ID de la tarjeta al archivo server.go
~~~

4. Ir a: [http://localhost:4242/pay.html](http://localhost:4242/pay.html)

~~~
Le damos al botón "Pagar" para realizar cargos a la tarjeta del cliente
~~~

Toda la información se deberá ver reflejada en nuestro dashboard de openpay