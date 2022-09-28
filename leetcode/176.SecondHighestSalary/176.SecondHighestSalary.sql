-- inner select will return the highest 
-- outer select should return the second highest 

select max(salary) as SecondHighestSalary from employee 
where salary not in (select max(salary) from employee)