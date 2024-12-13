/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package com.mycompany.collection.book;

class Book {
    private String title;
    private int pages;
    private int price;

    public Book(String title, int pages, int price) {
        this.title = title;
        this.pages = pages;
        this.price = price;
    }

    public String getTitle() {
        return title;
    }

    public int getPages() {
        return pages;
    }

    public int getPrice() {
        return price;
    }

    @Override
    public String toString() {
        return title + " (" + pages + " page): " + price + " won";
    }
}