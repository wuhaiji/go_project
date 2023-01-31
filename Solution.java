import java.util.HashMap;

public class Solution {
    public static class Node {
        public Node pre;
        public Node next;
        public Integer key;
        public Integer val;

        public Node(Integer key, Integer val) {
            this.key = key;
            this.val = val;
        }

        public Node() {

        }
    }

    private final int capacity;
    private final Node head;
    private final Node tail;
    private int size;

    public Solution(int capacity) {
        this.capacity = capacity;
        head = new Node();
        tail = new Node();
        head.next = tail;
        tail.pre = head;
    }

    @Override
    public String toString() {
        if (head.next == tail) {
            return "{}";
        }
        Node node = head.next;
        StringBuilder sb = new StringBuilder();
        sb.append("{");
        for (; node.key != null; node = node.next) {
            sb
                    .append("\"").append(node.key).append("\"")
                    .append(":")
                    .append("\"").append(node.val).append("\"")
                    .append(",");
        }
        if (sb.length() > 1) {
            sb.deleteCharAt(sb.length() - 1);
            sb.deleteCharAt(sb.length() - 1);
        }
        sb.append(" ").append("}");
        return sb.toString();
    }

    public int get(int key) {
        int res = -1;
        Node node = map.get(key);
        if (node != null) {
            refreshNode(node);
            res = node.val;
        }
        return res;
    }

    private void refreshNode(Node node) {
        deleteNode(node);
        addLast(node);
    }

    private void deleteNode(Node node) {
        Node pre = node.pre;
        Node next = node.next;

        pre.next = next;
        next.pre = pre;

        map.remove(node.key);
    }

    private void addLast(Node node) {
        Node pre = tail.pre;

        pre.next = node;
        node.pre = pre;

        node.next = tail;
        tail.pre = node;

        map.put(node.key, node);
    }

    public final HashMap<Integer, Node> map = new HashMap<>();

    public void set(int key, int val) {
        Node node = map.get(key);
        if (node != null) {
            node.val = val;
            refreshNode(node);
        } else {
            node = new Node(key, val);
            size++;
            addLast(node);
            if (size > capacity) {
                deleteNode(head.next);
                size--;
            }
        }
    }
}

