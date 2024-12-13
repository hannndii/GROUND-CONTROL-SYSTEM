/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package soal;

import java.util.Scanner;

/**
 *
 * @author tupog
 */
public class Soal {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        int A = 0, J = 0, M = 0;
        Scanner input = new Scanner(System.in);
        System.out.println("How much your stock: ");
        int sumFruit = input.nextInt();
        input.nextLine();
        for(int i = 0; i < sumFruit; i++){
            boolean validInput = false;
            
            while(!validInput){
            System.out.println("Input name of fruit: ");
            String fruit = input.nextLine();
            if(fruit.equals("Apel")){
                A++;
                validInput = true;
            } else if(fruit.equals("Jeruk")){
                J++;
                validInput = true;
            } else if(fruit.equals("Mangga")){
                M++;
                validInput = true;
            } else {
                System.out.println("Masukkan tidak sesuai!");
            }
        }
        }
        System.out.println("Sum Of Apple: " + A);
        System.out.println("Sum Of Orange: " + J);
        System.out.println("Sum Of Mango: " + M);
    }
}
