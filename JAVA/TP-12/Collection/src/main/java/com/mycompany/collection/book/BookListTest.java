/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package com.mycompany.collection.book;
import java.util.ArrayList;
/**
 *
 * @author tupog
 */
public class BookListTest {

    public static void main(String[] args) {
        ArrayList<Book> books = new ArrayList<>();

        books.add(new Book("Rules of Java", 346, 15000));
        books.add(new Book("Oracle 11g", 560, 18000));
        books.add(new Book("Copying JSP", 271, 12000));

        for (int i = 0; i < books.size(); i++) {
            System.out.println((i + 1) + ". " + books.get(i));
        }
    }
}
