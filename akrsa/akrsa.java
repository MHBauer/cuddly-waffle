import java.math.BigInteger;
import java.util.Random;
class akrsa{
    public static void main(String[] args) {
        var random = new Random();
        var e = new BigInteger(4096, random);
        var d = new BigInteger(4096, random);
        var m = new BigInteger(4096, random);
        var x = new BigInteger(4096, random);
        var r = e.multiply(d).mod(m);
        var c = e.multiply(x).mod(m);
        var y1 = c.multiply(d).mod(m);
        var y2 = r.multiply(x).mod(m);
            
        System.out.println(y1);
        System.out.println(y2);
        System.out.println(y1.equals(y2));        
    }
}
