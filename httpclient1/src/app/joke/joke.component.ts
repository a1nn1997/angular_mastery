import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
//import { JokeService } from 'src/services/joke.service';

@Component({
  selector: 'app-joke',
  templateUrl: './joke.component.html',
  styleUrls: ['./joke.component.css']
})
export class JokeComponent implements OnInit {
  joke:string = 'Loading Joke';
  constructor(private http: HttpClient) { 
    
  }  //public also  //dependency injection
  ngOnInit(): void {
    this.fetchData()
  }
  fetchData(){
    this.http.get('https://api.chucknorris.io/jokes/random?category=political')
    //subscribe essential for data
   // this.jokeService.getJoke()
    .subscribe((data:any)=>{
     this.joke=data.value;
    })   
  }
}
