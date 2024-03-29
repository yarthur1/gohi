package main

import (
    "fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//func isSubStructure(A *TreeNode, B *TreeNode) bool {
//    return A != nil && B != nil && (recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)) // ||满足三者之一即可
//}
//
//func recur (A *TreeNode, B *TreeNode) bool {
//    if B == nil { // 说明越界匹配完成
//        return true
//    }
//    if A == nil || A.Val != B.Val {
//        return false
//    }
//    return recur(A.Left, B.Left) && recur(A.Right, B.Right) //判断A，B左右子树是否相等
//}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
   if B==nil {
       return false
   }
   if recur(A, B) {
       return true
   }
   return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func recur(A *TreeNode, B *TreeNode) bool {
   if A==nil && B==nil {
       return true
   }
   if A==nil || B==nil || A.Val!=B.Val {
       return false
   }

   return recur(A.Left,B.Left)&&recur(A.Right,B.Right)
}

func main() {
    var A,B *TreeNode
    A=&TreeNode {Val:1}
    B=&TreeNode {Val:2}
    A.Left=B

    fmt.Printf("%v",isSubStructure(A, B))

}
