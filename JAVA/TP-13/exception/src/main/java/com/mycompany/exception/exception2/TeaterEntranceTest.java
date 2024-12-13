/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package com.mycompany.exception.exception2;
import java.util.ArrayList;
/**
 *
 * @author tupog
 */
public class TeaterEntranceTest {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        ArrayList<Person> people = new ArrayList<Person>();
        
        people.add(new Person("Hong-Gil-Dong", 23));
        people.add(new Person("Hong-Gil-Sun", 17));
        
        for(Person person : people){
            System.out.println(person.getName()+"Mr:");
            
            try{
                if(person.getAge() < 19){
                    throw new Exception("Dilarang Masuk!!");
                } else{
                    System.out.println("Silahkan Masuk:)");
                }
            } catch(Exception e){
                System.out.println(e.getMessage());
            }
        }
    }
    
}
