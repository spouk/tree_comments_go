Simple example tree comments release on Go
------------------------------------------
Use -
1) obtaining a list of comments,
from the table where there are 2 elements - ID, PID (respectively, ID unique comment identifier in the table, PID -> ID parent identifier of this comment)
2) creating a comment tree
3) output through the wrapper, or direct use of the function in templates when displaying the context

example output
==============

  
  ---[Output for html-template]-----------------------------------------------------------------------------------
<ul>
<li id='1_0'> ROOT 1 </li>
<ul>
<li id='3_1'> Simpletext </li>
<li id='4_1'> Simpletext </li>
<ul>
<li id='9_4'> Simpletext </li>
<ul>
<li id='10_9'> Simpletext </li>
<li id='11_9'> Simpletext </li>
<li id='12_9'> Simpletext </li>
</ul>
</ul>
<li id='7_1'> Simpletext </li>
<li id='8_1'> Simpletext </li>
</ul>
<li id='2_0'> ROOT 2 </li>
<ul>
<li id='5_2'> Simpletext </li>
<li id='6_2'> Simpletext </li>
</ul>
</ul>
--------------------------------------------------------------------------------------
