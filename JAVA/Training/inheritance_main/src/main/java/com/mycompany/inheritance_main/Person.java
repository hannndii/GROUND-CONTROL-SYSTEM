/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.inheritance_main;

/**
 *
 * @author tupog
 */
public class Person {
    private String name;
    private String phoneNumber;
    
   public Person(String name, String phoneNumber){
       this.name = name;
       this.phoneNumber = phoneNumber;
   }
   
   
   String getName(){
       return name;
   }
   
   public String getPhoneNumber(){
       return phoneNumber;
   }
    
    public void dataPerson(){
        System.out.println("DATA PERTAMA");
        System.out.println("Nama: ", getName());
    }
}
