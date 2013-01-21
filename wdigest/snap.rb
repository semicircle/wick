#!/usr/bin/ruby

#snap.rb
#1. weibo login.
#2. open file ./posts/last to find the latest weibo id.
#3. get home_timeline with parameter ":since" = id.
#4. save each post as ./post/id.json

require "weibo_2"
require "json"

#1. weibo login.
WeiboOAuth2::Config.api_key = '2121083357'
WeiboOAuth2::Config.api_secret = '2f25088263caf1726717ad7181cd3665'
WeiboOAuth2::Config.redirect_uri = 'http://what.what'
client = WeiboOAuth2::Client.new

token = client.get_token_from_hash({:access_token => '2.00PODBACtfqX_C57f0c823c0QueEOE', :expires_at => 1515305277 })

if !token.validated?
    puts 'token not validated.'
    exit
end

postpath = ARGV[0]

#2. open file ./posts/last to find the latest weibo id.%
latest = IO.read '%s/latest' % postpath

#puts latest

posts = client.statuses.home_timeline ({:since_id => latest, :count => 100})

posts[:statuses].each do | post | 
    #puts post.to_json
    #IO.write '#{postpath}/%d' % post[:id], post.to_json
    thepath = File.join(postpath, post[:id].to_s)
    IO.write thepath, post.to_json
end

latest = posts[:statuses][0][:id].to_s

IO.write '%s/latest' % postpath, latest

puts 'snap: get %d posts, lastest id: %s' % [posts[:statuses].length , latest]

