/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */

package com.mycompany.student;

import java.util.Scanner;

/**
 *
 * @author tupog
 */
import java.util.Scanner;

public class Student {

    public static void main(String[] args) {

        // Membuat array berisi 20 objek studentData
        studentData[] arrStudent = new studentData[2];
        Scanner input = new Scanner(System.in);
        
        System.out.println("Data Siswa SMA Bojongsoang!");

        // Inisialisasi dan pengisian data untuk setiap siswa
        for (int i = 0; i < arrStudent.length; i++) {
            arrStudent[i] = new studentData(); // Inisialisasi objek baru

            System.out.print("Masukkan nama siswa: ");
            arrStudent[i].studentName = input.nextLine();

            System.out.print("Masukkan ID siswa: ");
            arrStudent[i].studentId = input.nextLine();

            System.out.print("Masukkan nilai Bahasa Korea: ");
            arrStudent[i].korean = input.nextDouble();

            System.out.print("Masukkan nilai Bahasa Inggris: ");
            arrStudent[i].english = input.nextDouble();

            System.out.print("Masukkan nilai Matematika: ");
            arrStudent[i].math = input.nextDouble();
            
            input.nextLine(); // Membersihkan buffer setelah nextDouble()
        }

        // Menampilkan informasi setiap siswa
        for (int i = 0; i < arrStudent.length; i++) {
            arrStudent[i].printStudentInfo();
        }

        // Contoh penggunaan method untuk mengubah ID siswa
        System.out.print("Masukkan nomor siswa yang ingin diubah ID-nya: ");
        String id = input.nextLine();
        boolean found = false; 
        for (int i = 0; i < arrStudent.length; i++) {
            if(arrStudent[i].studentId.equals(id)){
                System.out.print("Masukkan ID Baru: ");
                String idChange = input.nextLine();
                arrStudent[i].changeStudentId(idChange);
                arrStudent[i].printStudentInfo();
                found = true; 
                break; 
            }
        }
        if (!found) {
            System.out.print("ID Tidak ditemukan!"); 
        }

}
}




