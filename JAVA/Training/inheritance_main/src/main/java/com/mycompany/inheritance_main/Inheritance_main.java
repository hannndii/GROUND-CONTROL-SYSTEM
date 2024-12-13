/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */

package com.mycompany.inheritance_main;

import newpackage.Student;

/**
 *
 * @author tupog
 */
public class Inheritance_main {

    public static void main(String[] args) {
        Person person = new Person("handi", "0812");
        Student student = new Student("handi", "0812", "Mahasiswa");
        Teacher teacher = new Teacher("handi", "0812", "Dosen RPL");
        
        
        System.out.println("==== INI ADALAH INHERITANCE ====");
        person.eat();
        student.study();
        teacher.teach();
        System.out.println("=================================");
        
        
    }
}
