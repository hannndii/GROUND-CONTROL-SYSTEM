/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.collection.student;

/**
 *
 * @author tupog
 */
public class Student {
    private String name;
    private int koreanScore;
    private int englishScore;
    private int mathScore;

    public Student(String name, int koreanScore, int englishScore, int mathScore) {
        this.name = name;
        this.koreanScore = koreanScore;
        this.englishScore = englishScore;
        this.mathScore = mathScore;
    }

    @Override
    public String toString() {
        return name + " (Bahasa Korea: " + koreanScore +
                ", Bahasa Inggris: " + englishScore +
                ", Matematika: " + mathScore + ")";
    }
}

