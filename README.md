Simple example tree comments release on Go
------------------------------------------
Use -
1) obtaining a list of comments,
from the table where there are 2 elements - ID, PID (respectively, ID unique comment identifier in the table, PID -> ID parent identifier of this comment)
2) creating a comment tree
3) output through the wrapper, or direct use of the function in templates when displaying the context

example output
==============

[COMMENTSTREE]2017/10/07 12:53:51 /home/spouk/go/src/simple/algoritms/example_tree_comments_go/search_tree.go:58: Нода успешно добавлена к корневой ноде
[COMMENTSTREE]2017/10/07 12:53:51 /home/spouk/go/src/simple/algoritms/example_tree_comments_go/search_tree.go:58: Нода успешно добавлена к корневой ноде
[COMMENTSTREE]2017/10/07 12:53:51 /home/spouk/go/src/simple/algoritms/example_tree_comments_go/search_tree.go:68: Новая нода `&{3 1 [] {3 1 Simpletext}}` успешно добавлена к своему `родителю`: `&{1 0 [0xc42008a2d0] {1 0 ROOT 1}}`
[COMMENTSTREE]2017/10/07 12:53:51 /home/spouk/go/src/simple/algoritms/example_tree_comments_go/search_tree.go:68: Новая нода `&{4 1 [] {4 1 Simpletext}}` успешно добавлена к своему `родителю`: `&{1 0 [0xc42008a2d0 0xc42008a460] {1 0 ROOT 1}}`
[COMMENTSTREE]2017/10/07 12:53:51 /home/spouk/go/src/simple/algoritms/example_tree_comments_go/search_tree.go:68: Новая нода `&{5 2 [] {5 2 Simpletext}}` успешно добавлена к своему `родителю`: `&{2 0 [0xc42008a5f0] {2 0 ROOT 2}}`
[COMMENTSTREE]2017/10/07 12:53:51 /home/spouk/go/src/simple/algoritms/example_tree_comments_go/search_tree.go:68: Новая нода `&{6 2 [] {6 2 Simpletext}}` успешно добавлена к своему `родителю`: `&{2 0 [0xc42008a5f0 0xc42008a870] {2 0 ROOT 2}}`
[COMMENTSTREE]2017/10/07 12:53:51 /home/spouk/go/src/simple/algoritms/example_tree_comments_go/search_tree.go:68: Новая нода `&{7 1 [] {7 1 Simpletext}}` успешно добавлена к своему `родителю`: `&{1 0 [0xc42008a2d0 0xc42008a460 0xc42008aaf0] {1 0 ROOT 1}}`
[COMMENTSTREE]2017/10/07 12:53:51 /home/spouk/go/src/simple/algoritms/example_tree_comments_go/search_tree.go:68: Новая нода `&{8 1 [] {8 1 Simpletext}}` успешно добавлена к своему `родителю`: `&{1 0 [0xc42008a2d0 0xc42008a460 0xc42008aaf0 0xc42008ac80] {1 0 ROOT 1}}`
[COMMENTSTREE]2017/10/07 12:53:51 /home/spouk/go/src/simple/algoritms/example_tree_comments_go/search_tree.go:68: Новая нода `&{9 4 [] {9 4 Simpletext}}` успешно добавлена к своему `родителю`: `&{4 1 [0xc42008ae10] {4 1 Simpletext}}`
[COMMENTSTREE]2017/10/07 12:53:51 /home/spouk/go/src/simple/algoritms/example_tree_comments_go/search_tree.go:68: Новая нода `&{10 9 [] {10 9 Simpletext}}` успешно добавлена к своему `родителю`: `&{9 4 [0xc42008b040] {9 4 Simpletext}}`
[COMMENTSTREE]2017/10/07 12:53:51 /home/spouk/go/src/simple/algoritms/example_tree_comments_go/search_tree.go:68: Новая нода `&{11 9 [] {11 9 Simpletext}}` успешно добавлена к своему `родителю`: `&{9 4 [0xc42008b040 0xc42008b2c0] {9 4 Simpletext}}`
[COMMENTSTREE]2017/10/07 12:53:51 /home/spouk/go/src/simple/algoritms/example_tree_comments_go/search_tree.go:68: Новая нода `&{12 9 [] {12 9 Simpletext}}` успешно добавлена к своему `родителю`: `&{9 4 [0xc42008b040 0xc42008b2c0 0xc42008b540] {9 4 Simpletext}}`
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
[COMMENTSTREE]2017/10/07 12:53:51 /home/spouk/go/src/simple/algoritms/example_tree_comments_go/search_tree.go:123: `&{1 0 [0xc42008a2d0 0xc42008a460 0xc42008aaf0 0xc42008ac80] {1 0 ROOT 1}}`
[COMMENTSTREE]2017/10/07 12:53:51 /home/spouk/go/src/simple/algoritms/example_tree_comments_go/search_tree.go:123: `&{2 0 [0xc42008a5f0 0xc4200

