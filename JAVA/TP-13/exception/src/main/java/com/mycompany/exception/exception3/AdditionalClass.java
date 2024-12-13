/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package com.mycompany.exception.exception3;
import java.util.ArrayList;
/**
 *
 * @author tupog
 */
public class AdditionalClass {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        ArrayList<Subject> subjects = new ArrayList<>();
        subjects.add(new Subject("Sejarah", 86));
        subjects.add(new Subject("Geografi", 65));
        subjects.add(new Subject("Biologi", 58));
        subjects.add(new Subject("Fisika", 76));

        System.out.println("Pelajaran yang perlu kelas tambahan");
        for (Subject subject : subjects) {
            try {
                if (subject.getScore() < 70) {
                    throw new Exception(subject.getName() + " (" + subject.getScore() + " poin)");
                }
            } catch (Exception e) {
                System.out.println(e.getMessage());
            }
        }
    }
}
