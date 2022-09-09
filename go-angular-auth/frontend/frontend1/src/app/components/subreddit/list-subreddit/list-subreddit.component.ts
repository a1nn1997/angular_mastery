import { Component, OnInit } from '@angular/core';
import { throwError } from 'rxjs';
import { SubredditModel } from '../subreddit-response';
import { SubredditService } from '../subreddit.service';

@Component({
  selector: 'app-list-subreddit',
  templateUrl: './list-subreddit.component.html',
  styleUrls: ['./list-subreddit.component.css']
})
export class ListSubredditComponent implements OnInit {

  subreddits: Array<SubredditModel>;
  constructor(private subredditService: SubredditService) { }

  ngOnInit() {
    this.subredditService.getAllSubreddits().subscribe(data => {
      this.subreddits = data;
    }, error => {
      throwError(error);
    })
  }
  

}
