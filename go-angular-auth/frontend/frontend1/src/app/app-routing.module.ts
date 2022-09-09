import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './components/home/home.component';
import { CreatePostComponent } from './components/post/create-post/create-post.component';
import { ViewPostComponent } from './components/post/view-post/view-post.component';
import { CreateSubredditComponent } from './components/subreddit/create-subreddit/create-subreddit.component';
import { ListSubredditComponent } from './components/subreddit/list-subreddit/list-subreddit.component';
import { EdituserComponent } from './components/user/edituser/edituser.component';
import { LoginComponent } from './components/user/login/login.component';
import { RegisterComponent } from './components/user/register/register.component';
import { UserProfileComponent } from './components/user/user-profile/user-profile.component';

const routes: Routes = [
   {path: '', component: HomeComponent},
   { path: 'view-post/:id', component: ViewPostComponent },
   { path: 'user-profile/:name', component: UserProfileComponent },

  {path: 'login', component: LoginComponent},
  {path: 'register', component: RegisterComponent},
  { path: 'list-subreddits', component: ListSubredditComponent },
  { path: 'create-post', component: CreatePostComponent },
  { path: 'create-subreddit', component: CreateSubredditComponent },
  {path: 'edituser', component: EdituserComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
