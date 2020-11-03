import java.io.*; 
import java.util.*; 
public class Main { 
    public static int map[][]; 
    public static boolean visit[][]; 
    public static int dist[][]; 
    public static final int[] dx = {-1,0,1,0}; 
    public static final int[] dy = {0,1,0,-1}; 
    public static int islandNum = 1; 
    public static int N; 
    public static int min = Integer.MAX_VALUE; 
    public static void main(String[] args) throws Exception{ 
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in)); 
        StringTokenizer st = new StringTokenizer(br.readLine()); 
        N = Integer.parseInt(st.nextToken()); 
        map = new int[N][N]; 
        visit = new boolean[N][N]; 
        
        for (int i = 0; i < N; i++) { 
            st = new StringTokenizer(br.readLine()); 
            for (int j = 0; j < N; j++) { 
                map[i][j] = Integer.parseInt(st.nextToken()); 
            } 
        } 
        
        for (int i = 0; i < N; i++) { 
            for (int j = 0; j <N ; j++) { 
                if(!visit[i][j] && map[i][j] != 0){ 
                    SetNum(i, j); 
                    islandNum++;
                } 
            } 
        } 
        dist = new int[N][N]; 
        for (int i = 0; i < N; i++) { 
            for (int j = 0; j < N; j++) { 
                dist[i][j] = Integer.MAX_VALUE; 
            } 
        }
        for (int i = 0; i < N; i++) { 
            for (int j = 0; j < N; j++) { 
                if(map[i][j] != 0) { 
                    FindLoad(map[i][j], i , j, 0); 
                } 
            } 
        }

        System.out.println(min); 
    } 

    public static boolean CheckIsland(int land, int nx, int ny) { 
        return map[nx][ny] == land; 
    } 

    public static boolean Check(int nx, int ny) { 
        return (nx < 0 || nx >= N || ny < 0 || ny >= N); 
    } 

    public static void FindLoad(int land, int row, int col, int cnt) { 
        if(min <= cnt) 
            return; 
        for (int i = 0; i < 4; i++) { 
            int nx = row + dx[i]; 
            int ny = col + dy[i]; 
            if(Check(nx,ny)) 
                continue; 
            if(dist[nx][ny] <= cnt+1) 
                continue; 
            if(map[nx][ny] == 0) { 
                dist[nx][ny] = cnt + 1; 
                FindLoad(land , nx, ny, cnt + 1); 
                continue; 
            } 
            if(CheckIsland(land,nx,ny)) 
                continue; 
            else { 
                min = Math.min(cnt, min); 
                return; 
            } 
        } 
    } 
    public static void SetNum(int row, int col) { 
        visit[row][col] = true; 
        map[row][col] = islandNum; 
        for (int i = 0; i < 4; i++) { 
            int nx = row + dx[i]; 
            int ny = col + dy[i]; 
            if(Check(nx,ny))
                continue; 
            if(map[nx][ny] == 0 || visit[nx][ny]) 
                continue; 
            SetNum(nx,ny); 
        } 
    } 
}

