# Strings
## Resumen
```
%       cualquier cantidad de caracteres (incluye 0)
_       exactamente 1 carácter
ILIKE   case-insensitive
ESCAPE  buscar % o _ literal

GROUP BY agrupa filas por columnas
HAVING filtra después del agrupamiento
FILTER permite condiciones dentro de agregaciones

ORDER BY col ASC|DESC
LIMIT N → cantidad de filas
OFFSET N → salto (paginación)
NULLS FIRST / LAST → control de NULLs

```

## Details

### Strings
id | title
---+-------------
1  | Atlas
2  | Han
3  | Chyan
4  | Cloud Atlas

SELECT title FROM movie WHERE title LIKE '%tlas';
→ Atlas, Cloud Atlas

SELECT title FROM movie WHERE title LIKE 'Atlas';
→ Atlas

SELECT title FROM movie WHERE title LIKE '%Atlas';
→ Atlas, Cloud Atlas

SELECT title FROM movie WHERE title LIKE '%an';
→ Han, Chyan

### _
id | value
---+--------
1  | cat
2  | bat
3  | car
4  | cut
5  | coat
6  | at

SELECT value FROM word WHERE value LIKE '_at';
→ cat, bat

SELECT value FROM word WHERE value LIKE 'c_t';
→ cat, cut

SELECT value FROM word WHERE value LIKE 'co_t';
→ coat

SELECT value FROM word WHERE value LIKE '__t';
→ cat, bat, cut

SELECT value FROM word WHERE value LIKE '___';
→ cat, bat, car, cut

### % + _
id | code
---+---------
1  | A1X
2  | A2Y
3  | B1X
4  | AAAX
5  | AX
6  | A123X

SELECT code FROM code WHERE code LIKE 'A_%';
→ A1X, A2Y, AAAX, AX, A123X

SELECT code FROM code WHERE code LIKE 'A__X';
→ A1X, A2Y, B1X (NO), AAAX (NO)

SELECT code FROM code WHERE code LIKE 'A%X';
→ A1X, AAAX, A123X

SELECT code FROM code WHERE code LIKE '_1X';
→ A1X, B1X


### mayúsculas
id | name
---+---------
1  | sam
2  | Sam
3  | SAM
4  | samantha

SELECT name FROM user_name WHERE name LIKE 'sam';
→ sam

SELECT name FROM user_name WHERE name ILIKE 'sam';
→ sam, Sam, SAM

SELECT name FROM user_name WHERE name ILIKE 'sam%';
→ sam, Sam, SAM, samantha


### ESCAPE
id | filename
---+-------------
1  | file_1.txt
2  | file_2.txt
3  | fileA.txt
4  | file%.txt

SELECT filename FROM files WHERE filename LIKE 'file\_%' ESCAPE '\';
→ file_1.txt, file_2.txt

SELECT filename FROM files WHERE filename LIKE 'file\%%' ESCAPE '\';
→ file%.txt

### 
id | customer | amount
---+----------+--------
1  | A        | 100
2  | A        | 200
3  | B        | 150
4  | B        | 50
5  | C        | 300

SELECT customer, SUM(amount)
FROM orders
GROUP BY customer;

→ A, 300
→ B, 200
→ C, 300

SELECT customer, COUNT(*)
FROM orders
GROUP BY customer;

→ A, 2
→ B, 2
→ C, 1

### GROUP BY (multiple columns)
id | customer | product | amount
---+----------+---------+--------
1  | A        | X       | 100
2  | A        | Y       | 200
3  | A        | X       | 50
4  | B        | X       | 150

SELECT customer, product, SUM(amount)
FROM sales
GROUP BY customer, product;

→ A, X, 150
→ A, Y, 200
→ B, X, 150

### GROUP BY (having)
id | customer | amount
---+----------+--------
1  | A        | 100
2  | A        | 200
3  | B        | 150
4  | B        | 50

SELECT customer, SUM(amount)
FROM orders
GROUP BY customer
HAVING SUM(amount) > 200;

→ A, 300

### GROUP BY (with expressions)
id | customer | amount
---+----------+--------
1  | A        | 100
2  | A        | 200
3  | B        | 150
4  | B        | 50

SELECT
  customer,
  SUM(amount),
  SUM(amount) * 0.1 AS tax
FROM orders
GROUP BY customer;

→ A, 300, 30
→ B, 200, 20

### GROUP BY (order by)
SELECT customer, SUM(amount) AS total
FROM orders
GROUP BY customer
ORDER BY total DESC;

→ A, 300
→ B, 200

### COUNT (with conditions)
id | customer | amount
---+----------+--------
1  | A        | 100
2  | A        | 200
3  | B        | 150
4  | B        | 50

SELECT
  customer,
  COUNT(*) AS total_orders,
  COUNT(*) FILTER (WHERE amount > 100) AS big_orders
FROM orders
GROUP BY customer;

→ A, 2, 1
→ B, 2, 1

### Distinct
id | customer
---+----------
1  | A
2  | A
3  | B

SELECT DISTINCT customer FROM orders;
→ A, B

### PAGINATION
SELECT * FROM orders
ORDER BY amount DESC
LIMIT 2 OFFSET 1;

### NULLS
id | amount
---+--------
1  | 100
2  | NULL
3  | 200
4  | NULL

SELECT * FROM payments
ORDER BY amount ASC NULLS FIRST;

→ (2, NULL)
→ (4, NULL)
→ (1, 100)
→ (3, 200)

SELECT * FROM payments
ORDER BY amount ASC NULLS LAST;

→ (1, 100)
→ (3, 200)
→ (2, NULL)
→ (4, NULL)

