import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { Todo } from 'src/app/Todo';

@Component({
  selector: 'app-todo-add',
  templateUrl: './todo-add.component.html',
  styleUrls: ['./todo-add.component.css']
})
export class TodoAddComponent implements OnInit {
  title: string;
  desc:string;
  sno:number=Math.random() *100;
  @Output() todoAdd: EventEmitter<Todo>=new EventEmitter();
  constructor() { 
    this.title=''
    this.desc=''
  }

  ngOnInit(): void {
  }
  onSubmit(){
    const todo = {
      sno: this.sno,
      title:this.title,
      desc:this.desc,
      active:true
    }
    this.todoAdd.emit(todo);
  }
}
