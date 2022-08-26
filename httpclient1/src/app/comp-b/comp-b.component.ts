import { Component, OnInit } from '@angular/core';
import { CounterService } from 'src/services/counter.service';

@Component({
  selector: 'app-comp-b',
  templateUrl: './comp-b.component.html',
  styleUrls: ['./comp-b.component.css'],
  providers:[CounterService]
})
export class CompBComponent implements OnInit {

  constructor(private c:CounterService) { }
  
  ngOnInit(): void {
  }
  
  showCounter(){
    return this.c.getCounter()
  }
  
  updateCounter(){
    return this.c.updateCounter()
  }

}
