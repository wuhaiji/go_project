import org.junit.Test;

/**
 * Your Solution object will be instantiated and called as such:
 * Solution solution = new Solution(capacity);
 * int output = solution.get(key);
 * solution.set(key,value);
 */
public class TestLRU {
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

