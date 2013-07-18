GoFaker
================

gofaker is a port of Faker from Ruby.

How to use
-----------------

::

   import "bitbucket.org/r_rudi/gofaker"

   faker, _ := gofaker.NewFaker("en_US")
   fmt.Println(faker.name.Phone)
   fmt.Println(faker.name.Last_name)
   fmt.Println(faker.name.First_name)

You can pass a locale string to NewFaker().

Dictonary
---------------

gofaker read dictionary data from "dict" directory and separated as
locale string.

::

  


go faker search under the locale directory at first, if not find, then
search from "base" directroy.


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

