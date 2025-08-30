package main

/* monitor SQL:
有一个hive表，表名是t，字段是a和b，a和b的值都是正整数（不存在null），且a和b的最大值都小于100。t的数据行数小于1000。

编程实现下面SQL的逻辑，输出SQL的结果：
select sum(rn) from (
select row_number() over(PARTITION BY case when t1.a is not null then t1.a else t2.b end) as rn
from t as t1 full join t as t2
on t1.a = t2.b and t1.b>10 and t2.a>10
) tmp
*/
func main() {

}
