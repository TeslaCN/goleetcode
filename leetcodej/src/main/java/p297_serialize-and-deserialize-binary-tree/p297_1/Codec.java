import java.util.LinkedList;
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
public class Codec {

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        if (root == null) {
            return "[null]";
        }
        LinkedList<TreeNode> q = new LinkedList<>();
        q.add(root);
        StringBuilder b = new StringBuilder("[");
        while (!q.isEmpty()) {
            TreeNode node = q.poll();
            if (node != null) {
                b.append(node.val).append(",");
                q.add(node.left);
                q.add(node.right);
            } else {
                b.append("null,");
            }
        }
        b.append("]");
        return b.toString();
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        String[] values = data.substring(1, data.length()-1).split(",");
        if (values.length == 0 || values.length == 1 && "null".equals(values[0])) {
            return null;
        }
        TreeNode root = new TreeNode(Integer.parseInt(values[0]));
        int cur = 1;
        LinkedList<TreeNode> q = new LinkedList<>();
        q.add(root);
        while (!q.isEmpty() && cur < values.length) {
            TreeNode node = q.poll();
            if (cur < values.length) {
                if (!"null".equals(values[cur])) {
                    node.left = new TreeNode(Integer.parseInt(values[cur]));
                    q.add(node.left);
                }
                cur++;
            }
            if (cur < values.length) {
                if (!"null".equals(values[cur])) {
                    node.right = new TreeNode(Integer.parseInt(values[cur]));
                    q.add(node.right);
                }
                cur++;
            }
        }
        return root;
    }
}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));

/*
执行用时： 17 ms , 在所有 Java 提交中击败了 73.87% 的用户
内存消耗： 41.1 MB , 在所有 Java 提交中击败了 28.57% 的用户
*/
