/*Variabel java:
 * int(4 byte) = 1024
 * long(8 byte) = 350000000000L (Huruf L besar menandakan long), menyimpan angka besar
 * double(8 byte) = 3.14159265359, digunakan ketika membutuhkan nilai akurasi yang tinggi dan presisi 
 * float(4 byte) = 2.718f (wajib ada huruf "f"), cocok untuk aplikasi di mana presisi tidak terlalu penting.
 */
public class trial {
    public static void main(String[] args) {
        int sisi1, sisi2, sisi3;
        int i = 0;
        sisi1 = 1;
        sisi2 = 2;
        sisi3 = 3;

        while (i != 5) {
            if (sisi1 + sisi2 > sisi3 && sisi1 + sisi3 > sisi2 && sisi2 + sisi3 > sisi1) {
                System.out.println("Segitiga");
            } else {
                System.out.println("Bukan Segitiga");
            }
            i++;
        }

    }
}
