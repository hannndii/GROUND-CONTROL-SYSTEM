/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */

package com.mycompany.trial1;

import java.util.Scanner;


/**
 *
 * @author tupog
 */
public class Trial1 {
    
    public static void main(String[] args) {
    
        NewClass nClass = new NewClass();
        Scanner input = new Scanner(System.in);
        
        nClass.english = 10;
        nClass.korean = 10;
        nClass.math = 10;
        nClass.studentName = "handi";
        nClass.studentId = "1030";
        String iD = "1000";
        
        nClass.calculateAverage();
        System.out.println("Hasil; " + nClass.calculateAverage());
        System.out.println("Data Mahasiswa!");
        nClass.printStudentInfo();
        
        System.out.println("Ganti ID");
        System.out.println("ID Sekarang: " + nClass.studentId);
        nClass.changeStudentId(iD);
        
    }
        
        
    }
