/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package soal2;

import java.util.Scanner;

/**
 *
 * @author tupog
 */
public class Soal2 {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        System.out.println("Masukkan nilai: ");
        double nilai = input.nextDouble();
        
        double result = 1.6;
        
        System.out.println("Nilai: " + (nilai*result) + " Kilometer");
        
    }
}
