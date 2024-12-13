/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.inheritance_main;

/**
 *
 * @author tupog
 */
public class Teacher extends Person {
    private String subject;
    
    public Teacher(String name, String phoneNumber, String subject){
        super(name, phoneNumber); 
        this.subject = subject;
    }
    
    public void teach(){
        System.out.println("Mengajar");
    }
}
