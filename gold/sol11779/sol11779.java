import java.io.*;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.PriorityQueue;
import java.util.StringTokenizer;

class Pair implements Comparable<Pair> {
    int node;
    long weight;

    Pair(int a, long b) {
        node = a;
        weight = b;
    }

    @Override
    public int compareTo(Pair o) {
        return Long.compare(this.weight, o.weight);
    }
}

public class Main {
    static ArrayList<Pair>[] edges = new ArrayList[1001];
    static ArrayList<Integer> routes = new ArrayList<>();
    static long[] dist = new long[1100];
    static int[] route = new int[1100];

    static int nCity, nEdge, src, dest;

    static int pI(String s) {
        return Integer.parseInt(s);
    }

    static void dijkstra(int src) {
        Arrays.fill(dist, Integer.MAX_VALUE);

        PriorityQueue<Pair> pq = new PriorityQueue<>();
        pq.add(new Pair(src, 0));
        dist[src] = 0;
        route[src] = 0;
        // SRC에서 각 점까지의 최단 거리 구하기
        while (!pq.isEmpty()) {
            int node = pq.peek().node;
            long weight = pq.poll().weight;

            for(Pair p : edges[node]){
                int nextNode = p.node;
                long nextWeight = p.weight + weight;

                if (dist[nextNode] > nextWeight) {
                    dist[nextNode] = nextWeight;
                    pq.add(new Pair(nextNode, nextWeight));
                    route[nextNode] = node;
                }
            }
        }

        // 경로 구하기
        int node = dest;
        while (node != 0) {
            routes.add(node);
            node = route[node];
        }
    }

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));
        StringTokenizer st = new StringTokenizer(br.readLine());

        nCity = pI(st.nextToken());
        nEdge = pI(new StringTokenizer(br.readLine()).nextToken());

        for(int i = 1; i <= nCity; i++){
            edges[i] = new ArrayList<>();
        }

        for (int i = 0; i < nEdge; i++) {
            st = new StringTokenizer(br.readLine());
            int a, b, c;

            a = pI(st.nextToken());
            b = pI(st.nextToken());
            c = pI(st.nextToken());

            edges[a].add(new Pair(b,c));
        }

        st = new StringTokenizer(br.readLine());
        src = pI(st.nextToken());
        dest = pI(st.nextToken());

        dijkstra(src);

        bw.write(dist[dest] + "\n");
        bw.write(routes.size() + "\n");
        for (int i = routes.size() - 1; i >= 0; i--)
            bw.write(routes.get(i) + " ");
        bw.flush();
    }
}