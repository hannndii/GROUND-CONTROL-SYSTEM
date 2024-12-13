/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package newpackage;

import com.mycompany.inheritance_main.Person;

/**
 *
 * @author tupog
 */
public class Student extends Person{
    private String departement;
    
    public Student(String name, String phoneNumber, String departement){
        super(name, phoneNumber);
        this.departement = departement;
    }
    
    public void study(){
        System.out.println("Belajar");
    }
    
}
