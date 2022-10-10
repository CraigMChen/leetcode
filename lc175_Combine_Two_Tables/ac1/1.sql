-- 由于 Person 表中并不是每个人都有 Address 信息，所以要用 left join
select FirstName, LastName, city, state
from Person left join Address on Person.PersonId = Address.PersonId