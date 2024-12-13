/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package com.mycompany.collection.student;
import java.util.HashMap;

/**
 *
 * @author tupog
 */
public class StudentTest {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        HashMap<String, Student> studentMap = new HashMap<>();

        studentMap.put("STU001", new Student("Lee Sun-shin", 80, 90, 95));
        studentMap.put("STU002", new Student("Kim Yu-shin", 95, 89, 92));
        studentMap.put("STU003", new Student("Kang Gam-chan", 88, 97, 94));

        for (String nim : studentMap.keySet()) {
            System.out.println(nim + " - " + studentMap.get(nim));
        }
    }
}