GoFaker
================

gofaker is a port of Faker from Ruby.

How to use
-----------------

::

   import (
   	"fmt"
   	"bitbucket.org/r_rudi/gofaker"
   )

   func main(){
   	faker, _ := gofaker.NewFaker("en_US", "src/bitbucket.org/r_rudi/gofaker/dict/")
   	fmt.Printf("Name: %s %s\n", faker.Name.First_name, faker.Name.Last_name)
   	fmt.Printf("Phone: %s\n", faker.Name.Phone)
   	fmt.Printf("%s %s %s, %s\n", faker.Address.Building_number, faker.Address.Street, faker.Address.City, faker.Address.State)
   	fmt.Printf("Zip: %s\n", faker.Address.Zipcode)
   }

Result::

   Name: Lois Evans
   Phone: 2-(078)367-1331
   3 Boyd Tustin, Ohio
   Zip: 65022-1364


Dictonary
---------------

gofaker read dictionary data from "dict" directory and separated as
locale string.

::

   dict
   |-- base
   |   |-- building_numbers
   |   |-- cities
   |   |-- company_names
   |   |-- phone
   |   `-- ...
   `-- ja_JP
       |-- building_numbers
       |-- female_first_names
       |-- last_names
       |-- male_first_names
       |-- phone
       `-- zipcode

gofaker searchs under the locale directory at first, if not find, then
search from "base" directroy.

alt
-----------

A dictonary file has one value by each line. But some files such as
last_names and first_names have multiple values which are separeted by "|". ex) california|CA

This is the alt value and can be used as "_alt" valiable.

::

  fmt.Println(faker.Name.Last_name_alt)

This can be used as "Furigana" or Abbrev value.

Available Data
------------------

name
++++++++

- First_name
- Lirst_name
- First_name_Alt (used for "Furigana", if dict has)
- Lirst_name_Alt
- Phone

address
++++++++

- Country
- City
- State
- Street
- Building_number
- Zipcode

company
++++++++

- Name
- Phone

internet
++++++++

- Url
- Domain
- Email

The giants on whose shoulders this works stands
----------------------------------------------------

- faker: http://faker.rubyforge.org/
- forgery: https://github.com/sevenwire/forgery

  - base data is copied from forgery. many thanks.

License
------------------

MIT License

