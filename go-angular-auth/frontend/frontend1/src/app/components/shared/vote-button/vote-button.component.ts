import { Component, Input, OnInit } from '@angular/core';
import { throwError } from 'rxjs';
import { PostModel } from '../post-model';
import { PostService } from '../post.service';
import { VoteService } from '../vote.service';
import { VotePayload } from './vote-payload';
import { VoteType } from './vote-type';
import { faArrowUp, faArrowDown } from '@fortawesome/free-solid-svg-icons';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-vote-button',
  templateUrl: './vote-button.component.html',
  styleUrls: ['./vote-button.component.css']
})
export class VoteButtonComponent implements OnInit {
@Input() post: PostModel;
  votePayload: VotePayload;
  faArrowUp = faArrowUp;
  faArrowDown = faArrowDown;
  upvoteColor: string;
  downvoteColor: string;
  isLoggedIn: boolean;

  constructor(private voteService: VoteService,
    
    private postService: PostService, private toastr: ToastrService) {

    this.votePayload = {
      voteType: undefined,
      postId: undefined
    }
  }

  ngOnInit(): void {
    this.updateVoteDetails();
  }

  upvotePost() {
    this.votePayload.voteType = String(VoteType.UPVOTE);
    this.vote();
    this.downvoteColor = '';
  }

  downvotePost() {
    this.votePayload.voteType = String(VoteType.DOWNVOTE);
    this.vote();
    this.upvoteColor = '';
  }

  private vote() {
    this.votePayload.postId = String(this.post.id);
    this.voteService.vote(this.votePayload).subscribe(() => {
      this.updateVoteDetails();
    }, error => {
      this.toastr.error(error.error.message);
      throwError(error);
    });
  }

  private updateVoteDetails() {
    this.postService.getPost(String(this.post.id)).subscribe(post => {
      this.post = post;
    });
  }

}
