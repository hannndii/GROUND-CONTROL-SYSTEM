/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package soal5;

import java.util.Scanner;

public class Soal5 {

    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        int[] nilai = new int[11];

        // Input nilai ke dalam array
        for (int i = 0; i < nilai.length; i++) {
            nilai[i] = input.nextInt();
        }
       
        int lastNumber = nilai[10];  // Menyimpan nilai elemen terakhir

        // Loop untuk membandingkan setiap elemen dengan elemen terakhir
        for (int i = 0; i < nilai.length-1; i++) {
            if (nilai[i] > lastNumber) {
                System.out.println(nilai[i] + " lebih besar dari " + lastNumber);
            } else if (nilai[i] < lastNumber) {
                System.out.println(nilai[i] + " lebih kecil dari " + lastNumber);
            } else {
                System.out.println(nilai[i] + " sama dengan " + lastNumber); 
            }
        }
    }
}

    

