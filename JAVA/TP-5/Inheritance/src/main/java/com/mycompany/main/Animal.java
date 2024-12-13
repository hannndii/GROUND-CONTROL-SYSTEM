/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.main;

/**
 *
 * @author tupog
 */
class Animal {
    // Attributes
    private int age;
    private String kind;
    
    // Constructor
    public Animal(int age, String kind) {
        this.age = age;
        this.kind = kind;
    }
    
    // Getters and Setters
    public int getAge() {
        return age;
    }
    
    public void setAge(int age) {
        this.age = age;
    }
    
    public String getKind() {
        return kind;
    }
    
    public void setKind(String kind) {
        this.kind = kind;
    }
    
    // Method
    public void run() {
        System.out.println("Lari!!!");
    }
    
    public void hunt() {
        System.out.println("Berburu!!");
    }
}

