import { Component, OnInit } from '@angular/core';
import { Todo } from 'src/app/Todo';
import {CdkDragDrop, moveItemInArray} from '@angular/cdk/drag-drop';

@Component({
  selector: 'app-todos',
  templateUrl: './todos.component.html',
  styleUrls: ['./todos.component.css']
})
export class TodosComponent implements OnInit {
  todos:Todo[];
  localItem: string | null;
  constructor() {
    this.todos = [];
    this.localItem = localStorage.getItem("todos")
    if(this.localItem==null){
      this.todos=[]
    }else{
      this.todos=JSON.parse(this.localItem);
    }
   }

  ngOnInit(): void {
  }
  deleteTodo(todo:Todo){
    this.todos.splice(this.todos.indexOf(todo),1);  //splice(i,1) will delete index at i with one entry
    localStorage.setItem("todos",JSON.stringify(this.todos));
  }
  addTodo(todo:Todo){
    this.todos.push(todo);  
    localStorage.setItem("todos",JSON.stringify(this.todos));  //storing in local storage to prevent default behaviour
  }
  toggleTodo(todo:Todo){
  this.todos[this.todos.indexOf(todo)].active=!this.todos[this.todos.indexOf(todo)].active;
    localStorage.setItem("todos",JSON.stringify(this.todos));  //storing in local storage to prevent default behaviour
  }
  drop(event: CdkDragDrop<string[]>) {
      moveItemInArray(this.todos, event.previousIndex, event.currentIndex);
    localStorage.setItem("todos",JSON.stringify(this.todos))
  }
}
