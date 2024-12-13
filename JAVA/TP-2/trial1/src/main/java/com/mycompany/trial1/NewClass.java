/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.trial1;

/**
 *
 * @author tupog
 */
public class NewClass {
    double korean, english, math, result;
    String studentName, studentId;
    
    // Method untuk menghitung rata-rata nilai
    public double calculateAverage() {
        result = (korean + english + math) / 3;
        return result;
    }

    // Method untuk mencetak informasi siswa
    public void printStudentInfo() {
        System.out.println("Nama Siswa: " + studentName);
        System.out.println("ID Siswa: " + studentId);
        System.out.println("Nilai Bahasa Korea: " + korean);
        System.out.println("Nilai Bahasa Inggris: " + english);
        System.out.println("Nilai Matematika: " + math);
        System.out.println("Rata-rata Nilai: " + calculateAverage());
        System.out.println("----------------------------");
    }
    
    // Method untuk mengubah ID siswa
    public void changeStudentId(String newId) {
        studentId = newId;
        System.out.println("ID Baru Siswa: " + newId);
    }
}



