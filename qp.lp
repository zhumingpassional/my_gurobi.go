\ Model qp
\ LP format - for model browsing. Use MPS format to capture full model detail.
Minimize
  2 x + [ 2 x ^2 + 2 x * y + 2 y ^2 + 2 y * z + 2 z ^2 ] / 2 
Subject To
 c0: x + 2 y + 3 z >= 4
 c1: x + y + z >= 1
Bounds
End
