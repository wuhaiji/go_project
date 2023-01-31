import org.junit.Test;

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

        node.next = null;
        node.pre = null;

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

    /**
     * Your Solution object will be instantiated and called as such:
     * Solution solution = new Solution(capacity);
     * int output = solution.get(key);
     * solution.set(key,value);
     */
    public static class TestLRU {
        @Test
        public void test1() {
            Solution s = new Solution(2);
            s.set(1, 1);
            System.out.println("set(1,1)" + s);

            s.set(2, 2);
            System.out.println("set(2,2)" + s);

            int i1 = s.get(1);
            System.out.println("get(1)  " + s);

            s.set(3, 3);
            System.out.println("set(3,3)" + s);

            int i2 = s.get(2);
            System.out.println("get(2)  " + s);

            s.set(4, 4);
            System.out.println("set(4,4)" + s);

            int i3 = s.get(1);
            System.out.println("get(1)  " + s);

            int i4 = s.get(3);
            System.out.println("get(3)  " + s);

            int i5 = s.get(4);
            System.out.println("get(4)  " + s);
        }

        @Test
        public void test2() {
            StringBuilder sb1 = new StringBuilder();
            Solution s = new Solution(3);

            sb1.append(set(s, 1, 1));
            sb1.append(set(s, 2, 2));
            sb1.append(set(s, 3, 3));
            sb1.append(get(s, 2));
            sb1.append(set(s, 4, 4));
            sb1.append(set(s, 5, 5));
            sb1.append(get(s, 2));

            System.out.println(sb1);

        }


        @Test
        public void test3() {
            StringBuilder sb1 = new StringBuilder();
            Solution s = new Solution(2);

            sb1.append(set(s, 1, 0));
            sb1.append(set(s, 2, 2));
            sb1.append(get(s, 1));
            sb1.append(set(s, 3, 3));
            sb1.append(get(s, 2));
            sb1.append(set(s, 4, 4 ));
            sb1.append(get(s, 1));
            sb1.append(get(s, 3));
            sb1.append(get(s, 4));

            System.out.println(sb1);


        }

        private static String get(Solution s, int x) {
            int i2 = s.get(x);
            System.out.println(s + ",map:" + s.map);

            return i2 + ",";
        }

        private static String set(Solution s, int key, int val) {
            s.set(key, val);
            System.out.println(s + ",map:" + s.map);
            return "null,";
        }
    }
}

